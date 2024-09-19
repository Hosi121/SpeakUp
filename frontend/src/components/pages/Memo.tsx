import { useState, useEffect } from "react";
import { Box, Button, Container, Stack, Typography } from "@mui/material";
import EditNoteIcon from "@mui/icons-material/EditNote";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import TopSection from "../utils/TopSection";
import { MemoInputField } from "../utils/MemoInputField";
import { fetchMemo, saveMemo } from "../../services/memoService";

const MemoContainer = () => {
  const [memo1, setMemo1] = useState("");
  const [memo2, setMemo2] = useState("");

  useEffect(() => {
    // ページ読み込み時にメモを取得
    const getMemo = async () => {
      try {
        const data = await fetchMemo();
        setMemo1(data.memo1 || "");
        setMemo2(data.memo2 || "");
      } catch (error) {
        console.error("Failed to fetch memos", error);
      }
    };
    getMemo();
  }, []);

  const handleSave = async () => {
    try {
      await saveMemo(memo1, memo2);
      console.log("Memos saved successfully");
      // 成功メッセージを表示するなどの処理
    } catch (error) {
      console.error("Failed to save memos", error);
    }
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
          <Typography variant="h4" sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}>
            <EditNoteIcon sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
            持ち込みメモ
          </Typography>
          <MemoInputField label="持ち込みメモ" value={memo1} onChange={setMemo1} />
          <Box sx={{ mt: 3 }}>
            <MemoInputField label="ワードリスト" value={memo2} onChange={setMemo2} />
          </Box>
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
