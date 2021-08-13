import { Box } from "@material-ui/core";
import { AppBar } from "../../components/AppBar";
import { MainContent } from "../../components/MainContent";
import { Page } from "../../components/Page";
import { PageHeading } from "../../components/PageHeading";

export function SettingsPage(): JSX.Element {
  return (
    <Box data-testid="SettingsPage">
      <Page>
        <AppBar />
        <MainContent>
          <PageHeading>Settings</PageHeading>
        </MainContent>
      </Page>
    </Box>
  );
}
