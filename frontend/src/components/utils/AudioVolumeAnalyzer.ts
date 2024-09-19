export class AudioVolumeAnalyzer {
  private audioContext: AudioContext;
  private analyser: AnalyserNode;
  private dataArray: Uint8Array;
  private source: MediaStreamAudioSourceNode | null = null;
  private animationFrameId: number | null = null;
  private setIsSpeak: (isSpeak: boolean) => void;

  constructor(setIsSpeak: (isSpeak: boolean) => void) {
    this.audioContext = new (window.AudioContext || (window as any).webkitAudioContext)();
    this.analyser = this.audioContext.createAnalyser();
    this.analyser.fftSize = 256;
    this.dataArray = new Uint8Array(this.analyser.frequencyBinCount);
    this.setIsSpeak = setIsSpeak;
  }

  start(stream: MediaStream): void {
    this.source = this.audioContext.createMediaStreamSource(stream);
    this.source.connect(this.analyser);
    this.updateVolume();
  }

  private updateVolume = () => {
    this.analyser.getByteFrequencyData(this.dataArray);
    const average = this.dataArray.reduce((acc, value) => acc + value, 0) / this.dataArray.length;
    const volume = Math.round((average / 255) * 100);
    this.setIsSpeak(volume > 10);

    this.animationFrameId = requestAnimationFrame(this.updateVolume);
  };

  stop(): void {
    if (this.source) {
      this.source.disconnect();
      this.source = null;
    }
    if (this.animationFrameId !== null) {
      cancelAnimationFrame(this.animationFrameId);
      this.animationFrameId = null;
    }
    this.audioContext.close();
  }
}
