<script>
  import router from "page";
  import Nav from "./components/Nav.svelte";
  import Foot from "./components/Foot.svelte";
  import Index from "./routes/Index.svelte";
  import User from "./routes/User.svelte";
  import Home from "./routes/Home.svelte";
  import NotFound from "./routes/NotFound.svelte";
  import Login from "./routes/Login.svelte";

  let page;
  let params;
  let user;

  router("/", () => {
    page = Index;
  });
  router("/home", () => {
    if (!user) {
      router.redirect("/");
    }
    page = Home;
  });
  router(
    "/user/:id",
    (ctx, next) => {
      params = ctx.params();
      next();
    },
    () => (page = User)
  );
  router("/login", () => (page = Login));
  router("/logout", () => (page = Login));
  router("/signup", () => (page = Login));
  router("/*", () => (page = NotFound));

  router.start();
</script>

<svelte:head>
  <title>Rolo</title>
</svelte:head>

<Nav {user} />
<main>
  <svelte:component this={page} {params} {user} />
</main>
<Foot />
