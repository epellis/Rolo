<script>
  export let user;
  console.log("User", user);
</script>

<nav class="navbar" role="navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">Rolo</a>
    <a class="navbar-item" href="/about">About</a>
    {#if user}
      <a class="navbar-item" href="/home">Home</a>
      <a class="navbar-item" href="/logout">Logout</a>
    {:else}
      <a class="navbar-item" href="/login">Login</a>
      <a class="navbar-item" href="/signup">Signup</a>
    {/if}
  </div>
</nav>
