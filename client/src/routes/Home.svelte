<script>
  import { userStore } from "../store.js";
  import CreatePost from "../components/CreatePost.svelte";

  let user;
  userStore.subscribe(newUser => {
    user = newUser;
  });
</script>

<section class="hero is-primary">
  <div class="hero-body">
    <h1 class="title">Welcome back, {user.username}!</h1>
    <h2 class="subtitle">Your saved links are below</h2>
  </div>
</section>
<section class="section">
  <div class="container">
    <CreatePost {userStore} />
  </div>
</section>
