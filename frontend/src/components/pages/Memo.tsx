import React, { useState } from "react";
import { TextField } from "@mui/material";
import { Box } from "@mui/system";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

const MemoContainer = () => {
  const [text, setText] = useState("");
  const maxLength = 500;

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setText(event.target.value.slice(0, maxLength));
  };

  return (
    <Box sx={{ margin: 2 }}>
      <TextField
        label="持ち込みメモ"
        multiline
        fullWidth
        variant="outlined"
        value={text}
        onChange={handleChange}
        helperText={`${text.length}/${maxLength}`}
        InputProps={{
          style: {
            height: "300px", // テキストエリアの高さを調整
          },
        }}
        rows={10}
      />
    </Box>
  );
};

export const Memo = () => {
  return (
    <BottomNavigationTemplate value="other">
      <MemoContainer />
    </BottomNavigationTemplate>
  );
};
