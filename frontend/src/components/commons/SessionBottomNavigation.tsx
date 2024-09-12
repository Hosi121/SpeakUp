import { Message, Mic, PriorityHigh, Search } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction, Box } from "@mui/material";

type sessionBottomNavigation = {
    value: number;
}

export const SessionBottomNavigation = ({ value }: sessionBottomNavigation) => {
    return (
        <Box>
            <BottomNavigation
                showLabels
                value={value}
            >
                <BottomNavigationAction icon={<Mic />} />
                <BottomNavigationAction icon={<Message />} />
                <BottomNavigationAction icon={<Search />} />
                <BottomNavigationAction icon={<PriorityHigh />} />
            </BottomNavigation>
        </Box>
    )
}
