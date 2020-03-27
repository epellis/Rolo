<script>
  import { userStore } from "../store.js";

  let user;
  userStore.subscribe(newUser => {
    user = newUser;
  });
</script>

<section class="hero is-primary">
  <div class="hero-body">
    <h1 class="title">Welcome back, {user.email}!</h1>
    <h2 class="subtitle">Your saved links are below</h2>
  </div>
</section>
