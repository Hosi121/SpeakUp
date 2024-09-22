import { Box, IconButton } from "@mui/material";
import SettingsIcon from "@mui/icons-material/Settings";
import NotificationModal from "../utils/NotificationModal";
import { useNavigate } from "react-router-dom";

const TopSection = () => {
  const navigate = useNavigate(); // Initialize useNavigate

  // Function to navigate to settings page
  const handleNavigateSettings = () => {
    navigate("/settings");
  };

  return (
    <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center", width: "100%" }}>
      {/* Notification Icon */}
      <Box>
        <NotificationModal />
      </Box>

      {/* Settings Icon (navigate to /settings) */}
      <IconButton onClick={handleNavigateSettings}>
        <SettingsIcon sx={{ fontSize: 40 }} />
      </IconButton>
    </Box>
  );
};

export default TopSection;
