<script>
  import { post } from "../utils.js";

  export let userStore;

  let url = "";
  let notes = "";
  let user;
  userStore.subscribe(newUser => {
    user = newUser;
  });

  async function submit(event) {
    let userid = user.userid;
    const response = await post("posts/create", { url, notes, userid });
    if (response.success) {
    }
  }
</script>

<div class="container">
  <div class="columns is-centered">
    <form class="box" on:submit|preventDefault={submit}>
      <div class="field">
        <h4 class="title is-4">Save a new link</h4>
      </div>
      <div class="field">
        <label for="" class="label">Url</label>
        <input type="text" class="input" bind:value={url} />
      </div>
      <div class="field">
        <label for="" class="label">Notes</label>
        <textarea class="textarea" bind:value={notes} />
      </div>
      <div class="field">
        <button class="button is-success" type="submit" disabled={!url}>
          Login
        </button>
      </div>
    </form>
  </div>
</div>
