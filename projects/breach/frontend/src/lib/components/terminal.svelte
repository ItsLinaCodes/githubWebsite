<script lang="ts">
  import { onMount, tick } from "svelte";
  import { colors } from "$lib/styles/colors";

  let inputRef: HTMLInputElement;

  type Line = { text: string; color: string };
  let lines: Line[] = [];
  let input = "";
  let outputEl: HTMLDivElement;
  let passwordMode = false;
  let prompt = "guest@corp-net:~$";
  let booting = true;
  let objective = "connect to the network";

  const objectives: [string, string][] = [
    ["ftp", "access the FTP server"],
    ["crack", "crack the password hash"],
    ["ssh", "move laterally via SSH"],
    ["exfil", "exfiltrate confidential.db"],
  ];

  async function appendLine(text: string, color: string) {
    if (color === "clear") {
      lines = [];
      return;
    }
    if (color === "update") {
      lines = [...lines.slice(0, -1), { text, color: colors.terminal.dim }];
    } else if (color === "password-prompt") {
      lines = [...lines, { text, color: colors.terminal.dim }];
      passwordMode = true;
    } else if (color === "green") {
      lines = [...lines, { text, color: colors.terminal.green }];
    } else if (color === "red") {
      lines = [...lines, { text, color: colors.terminal.red }];
    } else if (color === "white") {
      lines = [...lines, { text, color: colors.terminal.white }];
    } else if (color === "dim") {
      lines = [...lines, { text, color: colors.terminal.dim }];
    } else {
      lines = [...lines, { text, color }];
    }
    await tick();
    outputEl.scrollTop = outputEl.scrollHeight;
  }

  function handleEnter(e: KeyboardEvent) {
    if (e.key !== "Enter") return;
    if (passwordMode) {
      window.submitInput(input);
      passwordMode = false;
      input = "";
      return;
    }
    appendLine(`${prompt} ${input}`, colors.terminal.commandEcho);
    const cmd = input.trim().split(" ")[0];
    window.runCommand(input, appendLine);
    prompt = window.getPrompt();
    input = "";
    const next = objectives.find(([k]) => cmd === k);
    if (next) objective = next[1];
  }

  const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

  const logo = [
    "██████╗ ██████╗ ███████╗ █████╗  ██████╗██╗  ██╗",
    "██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔════╝██║  ██║",
    "██████╔╝██████╔╝█████╗  ███████║██║     ███████║",
    "██╔══██╗██╔══██╗██╔══╝  ██╔══██║██║     ██╔══██║",
    "██████╔╝██║  ██║███████╗██║  ██║╚██████╗██║  ██║",
    "╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝",
  ];

  const bootLines: [string, string][] = [
    ...logo.map((l): [string, string] => [l, "green"]),
    ["", "dim"],
    ["Penetration Testing Simulator  //  Educational Use Only", "dim"],
    ["", "dim"],
    ["────────────────────────────────────────────────────────", "dim"],
    ["", "dim"],
    ["[MISSION BRIEF]", "green"],
    ["Client:    MegaCorp Industries", "white"],
    ["Scope:     Internal network 10.0.0.0/24", "white"],
    ["Goal:      Locate and exfiltrate confidential.db", "white"],
    ["Rules:     Authorised pentest. Document all findings.", "dim"],
    ["", "dim"],
    ["You are connected to a VPN endpoint as: guest@corp-net", "dim"],
    ["Type 'help' to see available commands.", "dim"],
    ["", "dim"],
  ];

  onMount(async () => {
    inputRef.focus();
    for (const [text, color] of bootLines) {
      await appendLine(text, color);
      await sleep(80);
    }
    booting = false;
  });
</script>

<div
  class="flex flex-col w-full h-screen font-mono text-sm"
  style:background={colors.bg.primary}
  style:color={colors.terminal.green}
>
  <div
    class="flex items-center justify-between px-4 py-2 text-xs shrink-0"
    style:border-bottom="1px solid {colors.border.terminal}"
  >
    <span class="tracking-widest uppercase" style:color={colors.terminal.green}>
      BREACH v2.1 // PENETRATION TESTING SIMULATOR
    </span>
    <div class="flex items-center gap-4" style:color={colors.terminal.dim}>
      <span
        >objective: <span style:color={colors.terminal.green}>{objective}</span
        ></span
      >
      <a
        href="../../index.html"
        class="transition-colors hover:opacity-100"
        style:color={colors.terminal.dim}>→ portfolio</a
      >
    </div>
  </div>

  <div class="flex-1 overflow-y-auto px-6 py-4" bind:this={outputEl}>
    {#each lines as line}
      <div
        class="leading-snug whitespace-pre-wrap break-all"
        style:color={line.color}
      >
        {line.text}
      </div>
    {/each}
  </div>

  <div
    class="flex items-center px-6 py-3"
    style:border-top="1px solid {colors.border.terminal}"
  >
    <span class="shrink-0 mr-2" style:color={colors.terminal.green}
      >{booting ? "" : prompt}</span
    >
    <input
      type={passwordMode ? "password" : "text"}
      bind:value={input}
      on:keydown={handleEnter}
      bind:this={inputRef}
      disabled={booting}
      autocomplete="off"
      spellcheck="false"
      class="flex-1 bg-transparent border-none outline-none font-mono text-sm disabled:opacity-0"
      style:color={colors.terminal.green}
      style:caret-color={colors.terminal.green}
    />
  </div>
</div>
