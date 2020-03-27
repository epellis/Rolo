<script>
  import { userSession } from "../store.js";

  let user;
  userSession.subscribe(newSession => {
    user = newSession;
  });
</script>

<section class="hero is-primary">
  <div class="hero-body">
    <h1 class="title">Welcome back, {user.email}!</h1>
    <h2 class="subtitle">Your saved links are below</h2>
  </div>
</section>
