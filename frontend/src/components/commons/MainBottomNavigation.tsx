import { Home, LibraryBooks, Mic } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction } from "@mui/material";
import { useNavigate } from "react-router-dom";

type bottomNavigationProps = {
    value: number;
}

export const MainBottomNavigation = ({ value }: bottomNavigationProps) => {
    const navigate = useNavigate();
    return (
        <BottomNavigation
            showLabels
            value={value}
        >
            <BottomNavigationAction onClick={() => navigate("/record")} icon={<LibraryBooks />} />
            <BottomNavigationAction onClick={() => navigate("/home")} icon={<Home />} />
            <BottomNavigationAction onClick={() => navigate("/session")} icon={<Mic />} />
        </BottomNavigation>
    )
}
