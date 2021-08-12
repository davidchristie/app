import {
  Avatar,
  Box,
  Button,
  IconButton,
  Menu,
  MenuItem,
  RootRef,
} from "@material-ui/core";
import { MouseEventHandler } from "react";
import { useRef, useState } from "react";
import { useApi, User } from "../../api";

export interface UserMenuProps {
  user: User;
}

export function UserMenu({ user }: UserMenuProps): JSX.Element {
  const { useSignOut } = useApi();
  const signOut = useSignOut();
  const anchorEl = useRef<HTMLButtonElement | null>(null);
  const [open, setOpen] = useState(false);
  const toggleMenu: MouseEventHandler = () => setOpen((open) => !open);
  return (
    <Box data-testid="UserMenu" display="inline-block">
      <IconButton
        aria-controls="user-menu"
        aria-haspopup="true"
        data-testid="UserMenu__iconButton"
        onClick={toggleMenu}
        ref={anchorEl}
      >
        <Avatar src={user.avatarUrl} />
      </IconButton>
      <Menu
        anchorEl={anchorEl.current}
        anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
        data-testid="UserMenu__menu"
        getContentAnchorEl={null}
        id="user-menu"
        keepMounted
        onClick={toggleMenu}
        onClose={toggleMenu}
        open={open}
        transformOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <MenuItem data-testid="UserMenu__signOut" onClick={signOut}>
          Sign out
        </MenuItem>
      </Menu>
    </Box>
  );
}
