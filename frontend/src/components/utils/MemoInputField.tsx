import { TextField } from "@mui/material";
import { Dispatch, SetStateAction } from "react";

interface MemoInputFieldProps {
  value: string;
  setValue: Dispatch<SetStateAction<string>>;
  label: string;
  maxLength?: number;
  height?: string;
}

export const MemoInputField: React.FC<MemoInputFieldProps> = ({
  value,
  setValue,
  label,
  maxLength = 500,
  height = "300px",
}) => {
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const text = event.target.value.slice(0, maxLength);
    setValue(text);
  };

  return (
    <TextField
      label={label}
      multiline
      fullWidth
      variant="outlined"
      value={value}
      onChange={handleChange}
      helperText={`${value.length}/${maxLength}`}
      InputProps={{
        style: {
          height: height,
        },
      }}
      rows={10}
      sx={{ bgcolor: "secondary.main", borderRadius: 2 }}
    />
  );
};
