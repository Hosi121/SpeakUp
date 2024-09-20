import { Card, CardContent, Typography, List, ListItem, ListItemText, IconButton } from "@mui/material";
import topics from "../../mock/topics.json";
import { useEffect, useState } from "react";
import { Close } from "@mui/icons-material";

type TopicPopupProps = {
  isVisible: boolean;
  onClose: () => void;
};

export const TopicPopup = ({ isVisible, onClose }: TopicPopupProps) => {
  const [topicData, setTopicData] = useState<string[]>([]);

  useEffect(() => {
    const allTopics = topics.flatMap((topic) => topic.topics);
    setTopicData(allTopics);
  }, []);

  if (!isVisible) {
    return null;
  }
  return (
    <Card sx={{ width: "80vw", border: "2px solid #e91e63", borderRadius: "16px", p: 1, position: "fixed", bottom: "0", left: "50%", transform: "translate(-50%, -50%)" }}>
      <CardContent>
        <IconButton
          aria-label="close"
          onClick={onClose}
          sx={{
            position: "absolute",
            right: 8,
            top: 8,
            color: "#e91e63",
          }}
        >
          <Close />
        </IconButton>
        <Typography variant="h6" color="primary.main" gutterBottom textAlign="left" width="90%">
          まだ話していないトピックはありますか？
        </Typography>
        <List sx={{ height: "20vh" }}>
          {topicData.map((topic, index) => (
            <ListItem key={index} sx={{ height: "30%" }}>
              <ListItemText primary={topic} />
            </ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  );
};
