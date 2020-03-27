<script>
  import { links, Router } from "svelte-routing";

  export let user;
  export let url;
  console.log("User", user);
</script>

<nav class="navbar" role="navigation">
  <div class="navbar-brand" use:links>
    <Router {url}>
      <a class="navbar-item" href="/">Rolo</a>
      {#if user.isLoggedIn}
        <a class="navbar-item" href="/logout">Logout</a>
      {:else}
        <a class="navbar-item" href="/login">Login</a>
        <a class="navbar-item" href="/signup">Signup</a>
      {/if}
    </Router>
  </div>
</nav>
