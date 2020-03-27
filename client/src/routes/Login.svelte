<script>
  import { post } from "../utils.js";
  import { userStore } from "../store.js";
  import { navigate } from "svelte-routing";

  let email = "";
  let password = "";

  async function submit(event) {
    const response = await post("auth/login", { email, password });
    console.log("Response", response);
    if (response.success) {
      userStore.logIn(response);
      navigate("/");
    }
  }
</script>

<svelte:head>
  <title>Sign In</title>
</svelte:head>

<div class="container">
  <div class="columns is-centered">
    <form class="box" on:submit|preventDefault={submit}>
      <div class="field">
        <label for="" class="label">Email</label>
        <input type="email" class="input" bind:value={email} />
      </div>
      <div class="field">
        <label for="" class="label">Password</label>
        <input type="password" class="input" bind:value={password} />
      </div>
      <div class="field">
        <button
          class="button is-success"
          type="submit"
          disabled={!email || !password}>
          Login
        </button>
      </div>
    </form>
  </div>
</div>
