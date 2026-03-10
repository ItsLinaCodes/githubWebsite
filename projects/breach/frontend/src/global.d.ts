declare interface Window {
  Go: any;
  runCommand: (
    cmd: string,
    callback: (line: string, color: string) => void,
  ) => void;
}
