import { useState, useEffect } from "react";
import { Button, Container, Stack, Typography } from "@mui/material";
import EditNoteIcon from "@mui/icons-material/EditNote";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import TopSection from "../utils/TopSection";
import { MemoInputField } from "../utils/MemoInputField";
import api from "../../services/api";

const MemoContainer = () => {
  const [memo1, setMemo1] = useState("");
  const [memo2, setMemo2] = useState("");

  useEffect(() => {
    // Fetch memos on page load
    api
      .get("/memo")
      .then((response) => {
        setMemo1(response.data.memo1);
        setMemo2(response.data.memo2);
      })
      .catch((error) => {
        console.error("Failed to fetch memos", error);
      });
  }, []);

  const handleSave = () => {
    // Save memos on button click
    api
      .put("/memo", { memo1, memo2 })
      .then((response) => {
        console.log("Memos saved successfully");
      })
      .catch((error) => {
        console.error("Failed to save memos", error);
      });
  };

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100vh",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Typography
            variant="h4"
            sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}
          >
            <EditNoteIcon
              sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }}
            />
            持ち込みメモ
          </Typography>
          <MemoInputField
            memo1={memo1}
            memo2={memo2}
            onChangeMemo1={(e) => setMemo1(e.target.value)}
            onChangeMemo2={(e) => setMemo2(e.target.value)}
          />
          <Button
            variant="contained"
            fullWidth
            sx={{
              backgroundColor: "primary.main",
              color: "secondary.main",
              marginY: 2,
              "&:hover": {
                backgroundColor: "#FF3399",
              },
            }}
            onClick={handleSave}
          >
            保存
          </Button>
        </Stack>
      </Container>
    </Container>
  );
};

export const Memo = () => {
  return (
    <BottomNavigationTemplate value="other">
      <MemoContainer />
    </BottomNavigationTemplate>
  );
};