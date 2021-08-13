import { Avatar, Box, IconButton, Menu, MenuItem } from "@material-ui/core";
import { useRef, useState } from "react";
import { Link } from "react-router-dom";
import { useApi, User } from "../../api";

export interface UserMenuProps {
  user: User;
}

export function UserMenu({ user }: UserMenuProps): JSX.Element {
  const { useSignOut } = useApi();
  const signOut = useSignOut();
  const iconButtonRef = useRef<HTMLButtonElement | null>(null);
  const [open, setOpen] = useState(false);
  const toggleMenu = () => setOpen((open) => !open);
  return (
    <Box data-testid="UserMenu" display="inline-block">
      <IconButton
        aria-controls="user-menu"
        aria-haspopup="true"
        data-testid="UserMenu__iconButton"
        onClick={toggleMenu}
        ref={iconButtonRef}
      >
        <Avatar src={user.avatarUrl} />
      </IconButton>
      <Menu
        anchorEl={iconButtonRef.current}
        anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
        data-testid="UserMenu__menu"
        getContentAnchorEl={null}
        id="user-menu"
        keepMounted
        onClick={toggleMenu}
        open={open}
        transformOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <MenuItem
          component={Link}
          data-testid="UserMenu__settings"
          to="/settings"
        >
          Settings
        </MenuItem>
        <MenuItem data-testid="UserMenu__signOut" onClick={signOut}>
          Sign out
        </MenuItem>
      </Menu>
    </Box>
  );
}
