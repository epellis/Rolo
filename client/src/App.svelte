<script>
  import { userStore } from "./store.js";
  import { Router, Link, Route } from "svelte-routing";
  import Nav from "./components/Nav.svelte";
  import Foot from "./components/Foot.svelte";
  import Index from "./routes/Index.svelte";
  import Home from "./routes/Home.svelte";
  import Login from "./routes/Login.svelte";
  import Signup from "./routes/Signup.svelte";

  let user;
  userStore.subscribe(newUser => {
    user = newUser;
  });

  export let url = "";
</script>

<svelte:head>
  <title>Rolo</title>
</svelte:head>

<Nav {user} {url} />
<Router {url}>
  <div>
    <Route path="/">
      {#if !user.isLoggedIn}
        <Index />
      {:else}
        <Home />
      {/if}
    </Route>
    <Route path="/login">
      <Login />
    </Route>
    <Route path="/signup">
      <Signup />
    </Route>
  </div>
</Router>
