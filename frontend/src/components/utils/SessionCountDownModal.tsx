import React, { useEffect, useState } from "react";
import { Box, Typography } from "@mui/material";
import { Mic } from "@mui/icons-material";
import HomeLogo from "../../assets/homeLogo";

const SessionCountDownModal = () => {
  const [countdown, setCountdown] = useState(3);
  const [visible, setVisible] = useState(true);

  useEffect(() => {
    if (countdown > 0) {
      const timer = setTimeout(() => setCountdown(countdown - 1), 1000);
      return () => clearTimeout(timer);
    } else {
      setVisible(false);
    }
  }, [countdown]);

  if (!visible) {
    return null;
  }

  return (
    <Box
      sx={{
        height: "80vh",
        width: "90vw",
        bgcolor: "background.default",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "space-between",
      }}
    >
      <Box sx={{ textAlign: "center", mt: 4, width: "90%" }}>
        <Typography variant="h5" component="p" fontWeight="bold" textAlign="left">
          会話まで...
        </Typography>
        <Typography variant="h1" component="p" color="primary" fontWeight="bold" sx={{ mt: 2, fontSize: "6rem" }}>
          {countdown}
        </Typography>
      </Box>

      <Box sx={{ textAlign: "center", mb: 8 }}>
        <Typography variant="h4" color="primary" fontWeight="bold">
          Hello!と言って
        </Typography>
        <Typography variant="h4" color="primary" fontWeight="bold">
          自己紹介をして
        </Typography>
        <Typography variant="h4" color="primary" fontWeight="bold">
          始めよう！
        </Typography>
        <Box
          sx={{
            margin: "10px auto 0",
            width: "120px",
            height: "120px",
            backgroundColor: "secondary.main",
            borderRadius: 50,
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
          }}
        >
          <HomeLogo style={{ width: "80%" }} />
          <Mic fontSize="large" />
        </Box>
      </Box>
    </Box>
  );
};

export default SessionCountDownModal;
