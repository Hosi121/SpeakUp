import React, { useState, useEffect } from "react";
import { Accordion, AccordionSummary, AccordionDetails, Typography } from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import conversation_history from "../../mock/conversation_history.json"; // Import the JSON file directly

type SessionData = {
  date: string;
  previousDate: string;
  sessions: number;
  completionRate: string;
  comment: string;
  examples: string[];
};

const ConversationHistoryContainer = () => {
  const [data, setData] = useState<SessionData[]>([]);

  useEffect(() => {
    setData(conversation_history);
  }, []);

  return (
    <div>
      {data.map((item, index) => (
        <Accordion key={index}>
          <AccordionSummary expandIcon={<ExpandMoreIcon />}>
            <Typography>{item.date}</Typography>
          </AccordionSummary>
          <AccordionDetails>
            <Typography variant="body2">前回: {item.previousDate}</Typography>
            <Typography variant="body2">セッション数: {item.sessions}</Typography>
            <Typography variant="body2">達成率: {item.completionRate}</Typography>
            <Typography variant="body2">{item.comment}</Typography>
            <Typography variant="body2">参照可能な例: {item.examples.join(", ")}</Typography>
          </AccordionDetails>
        </Accordion>
      ))}
    </div>
  );
};

export const ConversationHistory = () => {
  return (
    <BottomNavigationTemplate value="other">
      <ConversationHistoryContainer />
    </BottomNavigationTemplate>
  );
};
