export type AppMode = "test" | "prod";

const resolveMode = (): AppMode => {
  const rawMode =
    import.meta.env.VITE_APP_MODE ??
    (import.meta.env.VITE_TEST_MODE === "true" ? "test" : "prod");
  const normalized =
    typeof rawMode === "string" ? rawMode.toLowerCase() : "prod";
  return normalized === "prod" || normalized === "production" ? "prod" : "test";
};

export const appMode: AppMode = resolveMode();
export const isTestMode = appMode === "test";
