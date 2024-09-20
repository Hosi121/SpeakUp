export const CircularLineSpinner = ({ size = 48, strokeWidth = 2, speed = 1 }) => {
  const lines = 12;
  const gapSize = 2;
  const radius = (size - strokeWidth) / 2;
  const center = size / 2;
  const color = "#000000";
  const transParentLines = 6;
  const calcStrokeColor = (num: number): string => {
    if (num < (lines - gapSize) - transParentLines) {
      return color;
    }
    const ratio = (num - ((lines - gapSize) - transParentLines)) / transParentLines;
    return `rgba(0, 0, 0, ${1 - ratio})`;
  }

  return (
    <div style={{ width: size, height: size }}>
      <svg
        width={size}
        height={size}
        viewBox={`0 0 ${size} ${size}`}
        xmlns="http://www.w3.org/2000/svg"
        style={{ animation: `spin ${1 / speed}s linear infinite` }}
      >
        {[...Array(lines - gapSize)].map((_, i) => {
          const angle = -(i * 360) / lines;
          const x1 = center + radius * Math.cos((angle * Math.PI) / 180);
          const y1 = center + radius * Math.sin((angle * Math.PI) / 180);
          const x2 = center + (radius - strokeWidth * 2) * Math.cos((angle * Math.PI) / 180);
          const y2 = center + (radius - strokeWidth * 2) * Math.sin((angle * Math.PI) / 180);
          return (
            <line
              key={i}
              x1={x1}
              y1={y1}
              x2={x2}
              y2={y2}
              stroke={calcStrokeColor(i)}
              strokeWidth={strokeWidth}
              strokeLinecap="round"
            />
          );
        })}
      </svg>
      <style>{`
        @keyframes spin {
          0% { transform: rotate(0deg); }
          100% { transform: rotate(360deg); }
        }
      `}</style>
    </div>
  );
};
