<script>
  import { links, Router, navigate } from "svelte-routing";

  export let url;
  export let userStore;

  let user;
  userStore.subscribe(newUser => {
    user = newUser;
  });

  function logout(event) {
    userStore.logOut();
    navigate("/");
  }
</script>

<nav class="navbar" role="navigation">
  <div class="navbar-brand" use:links>
    <Router {url}>
      <a class="navbar-item" href="/">Rolo</a>
      {#if user.isLoggedIn}
        <a class="navbar-item" href="/logout" on:click|preventDefault={logout}>
          Logout
        </a>
      {:else}
        <a class="navbar-item" href="/login">Login</a>
        <a class="navbar-item" href="/signup">Signup</a>
      {/if}
    </Router>
  </div>
</nav>
