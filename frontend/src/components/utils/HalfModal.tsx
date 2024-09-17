import React, { ReactNode } from "react";
import { Modal, Box, Typography, IconButton } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";

type HalfModalProps = {
  open: boolean;
  handleClose: () => void;
  children: ReactNode;
  title: string;
};

export const HalfModal = ({ open, handleClose, children, title }: HalfModalProps) => {
  return (
    <Modal open={open} onClose={handleClose} aria-labelledby="memo-modal-title">
      <Box
        sx={{
          position: "absolute",
          bottom: 0,
          left: 0,
          right: 0,
          bgcolor: "secondary.main",
          borderTopLeftRadius: 20,
          borderTopRightRadius: 20,
          boxShadow: 24,
          p: 3,
          height: "50%",
          overflow: "auto",
        }}
      >
        <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center", mb: 2 }}>
          <Typography variant="h6" component="h2" color="primary.main" fontWeight="bolder">
            {title}
          </Typography>
          <Box>
            <IconButton onClick={handleClose} aria-label="close">
              <CloseIcon />
            </IconButton>
          </Box>
        </Box>
        {children}
      </Box>
    </Modal>
  );
};
