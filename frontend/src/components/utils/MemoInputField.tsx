import { TextField } from '@mui/material';

interface MemoInputFieldProps {
  value: string;
  onChange: (value: string) => void;
  label: string;
  maxLength?: number;
}

export const MemoInputField: React.FC<MemoInputFieldProps> = ({
  value,
  onChange,
  label,
  maxLength = 500,
}) => {
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const text = event.target.value.slice(0, maxLength);
    onChange(text);
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
          height: '300px',
        },
      }}
      rows={10}
      sx={{ bgcolor: 'secondary.main', borderRadius: 2 }}
    />
  );
};