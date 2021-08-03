import { Box, Button } from "@material-ui/core";
import { FallbackProps } from "react-error-boundary";
import { MainContent } from "../../components/MainContent";
import { Page } from "../../components/Page";
import { PageHeading } from "../../components/PageHeading";

export function ErrorPage({ resetErrorBoundary }: FallbackProps): JSX.Element {
  return (
    <Box data-testid="ErrorPage">
      <Page>
        <MainContent>
          <PageHeading>Something went wrong</PageHeading>
          <Box marginTop={4}>
            <Button
              data-testid="ErrorPage__retryButton"
              onClick={resetErrorBoundary}
              variant="contained"
            >
              Retry
            </Button>
          </Box>
        </MainContent>
      </Page>
    </Box>
  );
}
