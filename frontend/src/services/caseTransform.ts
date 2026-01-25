const isPlainObject = (value: unknown): value is Record<string, unknown> => {
  if (Object.prototype.toString.call(value) !== "[object Object]") {
    return false;
  }
  const proto = Object.getPrototypeOf(value);
  return proto === Object.prototype || proto === null;
};

const toCamelKey = (key: string) => {
  const fromSnake = key.replace(/_([a-z0-9])/g, (_, char: string) =>
    char.toUpperCase()
  );
  return fromSnake.replace(/URL/g, "Url");
};

const toSnakeKey = (key: string) =>
  key.replace(/([A-Z])/g, "_$1").toLowerCase();

const transformKeys = (
  input: unknown,
  keyTransform: (key: string) => string
): unknown => {
  if (Array.isArray(input)) {
    return input.map((item) => transformKeys(item, keyTransform));
  }
  if (!isPlainObject(input)) {
    return input;
  }
  const output: Record<string, unknown> = {};
  Object.entries(input).forEach(([key, value]) => {
    output[keyTransform(key)] = transformKeys(value, keyTransform);
  });
  return output;
};

export const toCamelCase = (input: unknown): unknown =>
  transformKeys(input, toCamelKey);

export const toSnakeCase = (input: unknown): unknown =>
  transformKeys(input, toSnakeKey);
