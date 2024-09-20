import { Container } from "@mui/material";
import TopSection from "../utils/TopSection";
import { MicCheckScreen } from "../utils/MicCheckScreen";
import { useState } from "react";
import { MicCheckFinished } from "../utils/MicCheckFinished";

export const MicCheck = () => {
  const [isMicChecked, setIsMicChecked] = useState(false);
  const handleMicCheck = () => setIsMicChecked(true);
  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <Container sx={{ pt: 3, pb: 3 }}>
        <TopSection />
        {isMicChecked ? <MicCheckFinished /> : <MicCheckScreen isMicChecked={handleMicCheck} />}
      </Container>
    </Container>
  );
};
