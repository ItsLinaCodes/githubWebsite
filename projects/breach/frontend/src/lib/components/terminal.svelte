<script lang="ts">
  import { colors } from "$lib/styles/colors";
  import { onMount } from "svelte";

  let inputRef: HTMLInputElement;

  type Line = { text: string; color: string };
  let lines: Line[] = [];
  let input = "";
  let outputEl: HTMLDivElement;

  function appendLine(text: string, color: string) {
    lines = [...lines, { text, color }];
  }

  function handleEnter(e: KeyboardEvent) {
    if (e.key !== "Enter") return;
    appendLine(`$ ${input}`, colors.terminal.dim);
    window.runCommand(input, appendLine);
    input = "";
  }

  onMount(() => {
    inputRef.focus();
  });
</script>

<div class="flex flex-col w-full h-screen p-4 bg-[#0d0d0d] font-mono text-sm">
  <div
    class="flex flex-col-reverse flex-1 overflow-y-auto pb-2 scrollbar-thin scrollbar-thumb-[#555]"
    bind:this={outputEl}
  >
    {#each lines as line}
      <div
        class="leading-relaxed whitespace-pre-wrap break-all"
        style:color={line.color}
      >
        {line.text}
      </div>
    {/each}
  </div>

  <div
    class="flex items-center border-t border-[#555] pt-2"
    style:color={colors.terminal.green}
  >
    <span class="shrink-0">$&nbsp;</span>
    <input
      type="text"
      bind:value={input}
      on:keydown={handleEnter}
      bind:this={inputRef}
      autocomplete="off"
      spellcheck="false"
      class="flex-1 bg-transparent border-none outline-none font-mono text-sm caret-[#00ff41]"
      style:color={colors.terminal.green}
    />
  </div>
</div>
