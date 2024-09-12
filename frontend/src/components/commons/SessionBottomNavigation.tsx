import { Message, Mic, PriorityHigh, Search } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";

type sessionBottomNavigation = {
    value: number;
}

export const SessionBottomNavigation = ({ value }: sessionBottomNavigation) => {
    const navigate = useNavigate();
    return (
        <Box>
            <BottomNavigation
                showLabels
                value={value}
            >
                <BottomNavigationAction onClick={() => navigate("")} icon={<Mic />} />
                <BottomNavigationAction onClick={() => navigate("")} icon={<Message />} />
                <BottomNavigationAction onClick={() => navigate("")} icon={<Search />} />
                <BottomNavigationAction onClick={() => navigate("")} icon={<PriorityHigh />} />
            </BottomNavigation>
        </Box>
    )
}
