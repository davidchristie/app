import { Backdrop, CircularProgress } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";

interface Props {
  open: boolean;
}

const useStyles = makeStyles((theme) => ({
  backdrop: {
    background: theme.palette.background.default,
  },
}));

export function ProgressBackdrop({ open }: Props): JSX.Element {
  const classes = useStyles();
  return (
    <Backdrop
      className={classes.backdrop}
      data-testid="ProgressBackdrop"
      open={open}
    >
      <CircularProgress color="inherit" size={100} />
    </Backdrop>
  );
}
