<script>
  import { post } from "../utils.js";

  let url = "";
  let notes = "";

  async function submit(event) {
    const response = await post("posts/create", { url, notes });
    if (response.success) {
    }
  }
</script>

<div class="container">
  <div class="columns is-centered">
    <form class="box" on:submit|preventDefault={submit}>
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
