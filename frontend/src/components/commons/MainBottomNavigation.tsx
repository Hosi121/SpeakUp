import { Home, LibraryBooks, Mic } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction } from "@mui/material";

type bottomNavigationProps = {
    value: number;
}

export const MainBottomNavigation = ({ value }: bottomNavigationProps) => {
    return (
        <BottomNavigation
            showLabels
            value={value}
        >
            <BottomNavigationAction icon={<LibraryBooks />} />
            <BottomNavigationAction icon={<Home />} />
            <BottomNavigationAction icon={<Mic />} />
        </BottomNavigation>
    )
}
