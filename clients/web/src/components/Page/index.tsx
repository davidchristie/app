import { Box } from "@material-ui/core";
import { ReactNode } from "react";

export interface PageProps {
  children?: ReactNode;
}

export function Page({ children }: PageProps): JSX.Element {
  return (
    <Box height="100vh" data-testid="Page">
      {children}
    </Box>
  );
}
