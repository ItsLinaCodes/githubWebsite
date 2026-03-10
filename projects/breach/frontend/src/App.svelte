<script>
  import { onMount } from "svelte";
  import Terminal from "$lib/components/terminal.svelte";

  onMount(async () => {
    const go = new window.Go();
    const result = await WebAssembly.instantiateStreaming(
      fetch("/main.wasm"),
      go.importObject,
    );
    go.run(result.instance);
    console.log("WASM loaded");
  });
</script>

<main>
  <Terminal />
</main>
