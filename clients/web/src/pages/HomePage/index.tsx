import { Box } from "@material-ui/core";
import { AppBar } from "../../components/AppBar";
import { MainContent } from "../../components/MainContent";
import { Page } from "../../components/Page";
import { PageHeading } from "../../components/PageHeading";

export function HomePage(): JSX.Element {
  return (
    <Box data-testid="HomePage">
      <Page>
        <AppBar />
        <MainContent>
          <PageHeading>Home</PageHeading>
        </MainContent>
      </Page>
    </Box>
  );
}
