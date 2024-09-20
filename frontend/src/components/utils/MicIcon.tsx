import { Stack } from "@mui/material";
import HomeLogo from "../../assets/homeLogo";
import { Mic } from "@mui/icons-material";

export const MicIcon = () => {
  return (
    <Stack sx={{ width: "150px", height: "150px", borderRadius: 50, backgroundColor: "secondary.main", display: "flex", flexDirection: "column", justifyContent: "flex-end", alignItems: "center", p: 3 }}>
      <HomeLogo style={{ width: "100%", height: "fit-content", marginBottom: "10px" }} />
      <Mic sx={{ fontSize: "60px" }} />
    </Stack>
  );
};
