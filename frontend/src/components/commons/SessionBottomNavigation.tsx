import { Message, Mic, MicOff, PriorityHigh, Search } from "@mui/icons-material";
import { BottomNavigation, BottomNavigationAction, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";

type sessionBottomNavigation = ("wordList" | "assistance" | "topic")
type sessionBottomNavigationProps = {
    value: sessionBottomNavigation;
    isMute: boolean;
}


export const SessionBottomNavigation = ({ value, isMute }: sessionBottomNavigationProps) => {
    const navigate = useNavigate();
    const sessionBottomNavigationToIndex = (value: sessionBottomNavigation): number => {
        switch (value) {
            case "wordList":
                return 0;
            case "assistance":
                return 1;
            case "topic":
                return 2;
            default:
                return -1;
        }
    }
    const index = sessionBottomNavigationToIndex(value);
    return (
        <Box>
            <BottomNavigation
                showLabels
                value={index}
            >
                <BottomNavigationAction onClick={() => navigate("")} icon={isMute ? <MicOff /> : <Mic />} />
                <BottomNavigationAction onClick={() => navigate("")} icon={<Message />} />
                <BottomNavigationAction onClick={() => navigate("")} icon={<Search />} />
                <BottomNavigationAction onClick={() => navigate("")} icon={<PriorityHigh />} />
            </BottomNavigation>
        </Box>
    )
}
