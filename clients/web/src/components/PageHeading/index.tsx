import { Typography } from "@material-ui/core";
import { ReactNode } from "react";

export interface PageHeadingProps {
  children?: ReactNode;
}

export function PageHeading({ children }: PageHeadingProps): JSX.Element {
  return (
    <Typography data-testid="PageHeading" variant="h2">
      {children}
    </Typography>
  );
}
