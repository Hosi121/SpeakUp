import { Button, Icon } from "@mui/material";
import { ReactNode } from "react";
import { useNavigate } from "react-router-dom";

type IconButtonProps = {
  icon: ReactNode;
  text: string;
  url: string;
};

export const IconButton = ({ icon, text, url }: IconButtonProps) => {
  const navigate = useNavigate();
  return (
    <Button
      variant="contained"
      sx={{
        margin: "10px 0",
        height: "150px",
        width: "45%",
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        borderRadius: "10px",
        backgroundColor: "#f2d66f",
        color: "#da0063",
        fontSize: "1rem",
      }}
      onClick={() => navigate("/" + url)}
    >
      <Icon sx={{ width: "fit-content", height: "fit-content", color: "#000" }}>{icon}</Icon>
      {text}
    </Button>
  );
};
