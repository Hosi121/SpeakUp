import { Home, LibraryBooks, Mic } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction } from "@mui/material";
import { useNavigate } from "react-router-dom";

type bottomNavigationProps = {
  value: mainBottomNavigation;
};

export type mainBottomNavigation = "record" | "home" | "session" | "other";

export const MainBottomNavigation = ({ value }: bottomNavigationProps) => {
  const navigate = useNavigate();
  const mainBottomNavigationToIndex = (value: mainBottomNavigation): number => {
    switch (value) {
      case "record":
        return 0;
      case "home":
        return 1;
      case "session":
        return 2;
      default:
        return -1;
    }
  };
  const index = mainBottomNavigationToIndex(value);

  return (
    <BottomNavigation
      showLabels
      value={index}
      sx={{
        padding: "20px 0",
        height: "100px",
        position: "sticky",
        bottom: 0,
        zIndex: 100,
        backgroundColor: `${Theme.palette.background.default}`,
        boxShadow: "0 0 5px rgba(0, 0, 0, 0.6)",
      }}
    >
      <BottomNavigationAction
        onClick={() => navigate("/record")}
        icon={<LibraryBooks fontSize="large" />}
      />
      <BottomNavigationAction
        onClick={() => navigate("/home")}
        icon={<Home fontSize="large" />}
      />
      <BottomNavigationAction
        onClick={() => navigate("/session")}
        icon={<Mic fontSize="large" />}
      />
    </BottomNavigation>
  );
};
