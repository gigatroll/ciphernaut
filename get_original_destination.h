
// used for IPv4 sockets
int get_original_destination_4(int fd, char **orig_destination, int *port, int *err);

// used for IPv6 sockets
int get_original_destination_6(int fd, char *orig_destination, int *port);
