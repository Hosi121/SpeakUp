import { useState, useEffect, useRef } from "react";
import { Box, Typography } from "@mui/material";
import { styled } from "@mui/system";

const VisualizerContainer = styled(Box)(() => ({
  width: "140px",
  height: 25,
  display: "flex",
  alignItems: "flex-end",
  justifyContent: "center",
}));

const Bar = styled(Box)(() => ({
  width: 5,
  maxHeight: 25,
  backgroundColor: "#000",
  transition: "height 0.1s ease",
}));

const AudioVisualizer: React.FC = () => {
  const [audioData, setAudioData] = useState<Uint8Array>(new Uint8Array(128));
  const [error, setError] = useState<string | null>(null);

  const audioContextRef = useRef<AudioContext | null>(null);
  const analyserRef = useRef<AnalyserNode | null>(null);
  const dataArrayRef = useRef<Uint8Array | null>(null);
  const sourceRef = useRef<MediaStreamAudioSourceNode | null>(null);
  const rafIdRef = useRef<number | null>(null);

  useEffect(() => {
    startListening();

    return () => {
      if (rafIdRef.current) {
        cancelAnimationFrame(rafIdRef.current);
      }
      if (audioContextRef.current) {
        audioContextRef.current.close();
      }
    };
  }, []);

  const compressArray = (array: Uint8Array): Uint8Array => {
    const length = array.length * 0.7; // 高周波はカット
    return new Uint8Array(10).map((_, i) => {
      // startからendまでの平均値を取得
      const start = Math.floor(i * (length / 10));
      const end = Math.floor((i + 1) * (length / 10));
      const sum = array.slice(start, end).reduce((acc, val) => acc + val, 0);
      const average = sum / (end - start);
      const max = 255 * 1.0; // 表示上の頭打ち振幅を100%にする
      const min = 255 * 0.05; // 5%までは無音扱い
      const actualValue = Math.max(0, Math.min(255, average));
      const height = 25;
      return Math.round((actualValue / (max - min)) * height);
    });
  };

  const startListening = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });

      audioContextRef.current = new AudioContext();
      analyserRef.current = audioContextRef.current.createAnalyser();
      dataArrayRef.current = new Uint8Array(analyserRef.current.frequencyBinCount);
      sourceRef.current = audioContextRef.current.createMediaStreamSource(stream);
      sourceRef.current.connect(analyserRef.current);

      const updateAudioData = () => {
        if (analyserRef.current && dataArrayRef.current) {
          analyserRef.current.getByteFrequencyData(dataArrayRef.current);
          const scaledData = compressArray(dataArrayRef.current);

          setAudioData(scaledData);
        }
        rafIdRef.current = requestAnimationFrame(updateAudioData);
      };

      updateAudioData();
    } catch (error) {
      console.error("Error accessing microphone:", error);
      setError("マイクへのアクセスが拒否されました。ブラウザの設定を確認してください。");
    }
  };

  return (
    <Box>
      {error ? (
        <Typography color="error" align="center" sx={{ mb: 2 }}>
          {error}
        </Typography>
      ) : (
        <Box sx={{ mb: 2 }}>
          <VisualizerContainer sx={{ display: "flex", justifyContent: "space-between" }}>
            {Array.from(audioData.slice(0, 10)).map((value, index) => (
              <Bar
                key={index}
                style={{
                  height: `${Math.min(value, 25)}px`,
                }}
              />
            ))}
          </VisualizerContainer>
          <VisualizerContainer sx={{ display: "flex", justifyContent: "space-between", transform: "scaleY(-1)" }}>
            {Array.from(audioData.slice(0, 10)).map((value, index) => (
              <Bar
                key={index}
                style={{
                  height: `${Math.min(value, 25)}px`,
                }}
              />
            ))}
          </VisualizerContainer>
        </Box>
      )}
    </Box>
  );
};

export default AudioVisualizer;
