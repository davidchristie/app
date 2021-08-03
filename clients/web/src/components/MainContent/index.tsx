import { Box, Container } from "@material-ui/core";
import { ReactNode } from "react";

export interface MainContentProps {
  children?: ReactNode;
}

export function MainContent({ children }: MainContentProps): JSX.Element {
  return (
    <Container data-testid="MainContent">
      <Box marginTop={4}>{children}</Box>
    </Container>
  );
}
