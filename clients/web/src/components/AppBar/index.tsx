import {
  AppBar as MuiAppBar,
  Box,
  Button,
  createStyles,
  makeStyles,
  Toolbar,
  Typography,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import { useApi } from "../../api";

const useStyles = makeStyles((theme) =>
  createStyles({
    title: {
      color: "inherit",
      display: "inline-block",
      textDecoration: "inherit",
    },
    toolbar: {
      display: "flex",
      height: "100%",
      justifyContent: "space-between",
      width: "100%",
    },
  })
);

export function AppBar(): JSX.Element {
  const classes = useStyles();
  const { useSession } = useApi();
  const session = useSession();
  return (
    <MuiAppBar color="default" data-testid="AppBar" position="static">
      <Toolbar>
        <Box
          display="flex"
          height="100%"
          justifyContent="space-between"
          width="100%"
        >
          <Typography
            className={classes.title}
            component={Link}
            to="/"
            variant="h6"
          >
            App
          </Typography>
          <Box>
            {session.data !== undefined && (
              <>
                {session.data.user === null && (
                  <>
                    <Button
                      color="inherit"
                      component={Link}
                      data-testid="AppBar__signInButton"
                      to="/signin"
                      variant="contained"
                    >
                      Sign in
                    </Button>
                  </>
                )}
                {session.data.user !== null && (
                  <>
                    <Button
                      color="inherit"
                      component="a"
                      data-testid="AppBar__signOutButton"
                      href="/api/v1/auth/signout"
                    >
                      Sign out
                    </Button>
                  </>
                )}
              </>
            )}
          </Box>
        </Box>
      </Toolbar>
    </MuiAppBar>
  );
}
