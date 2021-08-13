import { Route, RouteProps, Switch } from "react-router-dom";
import { useApi } from "./api";
import { ProgressBackdrop } from "./components/ProgressBackdrop";
import { HomePage } from "./pages/HomePage";
import { NotFoundPage } from "./pages/NotFoundPage";
import { SettingsPage } from "./pages/SettingsPage";
import { SignInPage } from "./pages/SignInPage";

const authenticatedRoutes: RouteProps[] = [
  {
    exact: true,
    path: "/",
    component: HomePage,
  },
  {
    exact: true,
    path: "/settings",
    component: SettingsPage,
  },
  {
    path: "*",
    component: NotFoundPage,
  },
];

const unauthenticatedRoutes: RouteProps[] = [
  {
    exact: true,
    path: "/",
    component: HomePage,
  },
  {
    exact: true,
    path: "/signin",
    component: SignInPage,
  },
  {
    path: "*",
    component: NotFoundPage,
  },
];

export function Routes(): JSX.Element {
  const { useSession } = useApi();
  const session = useSession();
  if (session.data === undefined) {
    return <ProgressBackdrop open />;
  }
  const routes =
    session.data.user === null ? unauthenticatedRoutes : authenticatedRoutes;
  return (
    <Switch>
      {routes.map((route, index) => (
        <Route {...route} key={index} />
      ))}
    </Switch>
  );
}
