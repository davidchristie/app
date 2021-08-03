import { Box } from "@material-ui/core";
import { AppBar } from "../../components/AppBar";
import { MainContent } from "../../components/MainContent";
import { Page } from "../../components/Page";
import { PageHeading } from "../../components/PageHeading";

export function NotFoundPage(): JSX.Element {
  return (
    <Box data-testid="NotFoundPage">
      <Page>
        <AppBar />
        <MainContent>
          <PageHeading>Not Found</PageHeading>
        </MainContent>
      </Page>
    </Box>
  );
}
