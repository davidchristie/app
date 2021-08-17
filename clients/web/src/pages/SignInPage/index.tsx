import { Box, Button } from "@material-ui/core";
import { useApi } from "../../api";
import { AppBar } from "../../components/AppBar";
import { MainContent } from "../../components/MainContent";
import { Page } from "../../components/Page";
import { PageHeading } from "../../components/PageHeading";

export function SignInPage(): JSX.Element {
  const { useAuthorize } = useApi();
  const authorize = useAuthorize();
  return (
    <Box data-testid="SignInPage">
      <Page>
        <AppBar />
        <MainContent>
          <PageHeading>Sign In</PageHeading>
          <Box marginTop={4}>
            <Button
              component="a"
              data-testid="SignInPage__GitHubAuthorizeButton"
              onClick={() => authorize("github")}
              variant="contained"
            >
              Sign in with GitHub
            </Button>
          </Box>
          <Box marginTop={4}>
            <Button
              component="a"
              data-testid="SignInPage__GoogleAuthorizeButton"
              onClick={() => authorize("google")}
              variant="contained"
            >
              Sign in with Google
            </Button>
          </Box>
        </MainContent>
      </Page>
    </Box>
  );
}
