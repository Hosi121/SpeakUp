import { Home, LibraryBooks, Mic } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction } from "@mui/material";
import { useNavigate } from "react-router-dom";

type bottomNavigationProps = {
    value: mainBottomNavigation;
}

type mainBottomNavigation = "record" | "home" | "session" | "other";

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
    }
    const index = mainBottomNavigationToIndex(value);

    return (
        <BottomNavigation
            showLabels
            value={index}
        >
            <BottomNavigationAction onClick={() => navigate("/record")} icon={<LibraryBooks />} />
            <BottomNavigationAction onClick={() => navigate("/home")} icon={<Home />} />
            <BottomNavigationAction onClick={() => navigate("/session")} icon={<Mic />} />
        </BottomNavigation>
    )
}

