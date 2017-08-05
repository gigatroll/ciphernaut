#include <stdlib.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <linux/netfilter_ipv4.h> // SO_ORIGINAL_DST is defined here
#include <errno.h>
#include <strings.h>

#include <stdio.h>

int get_original_destination_4(int fd, char **orig_destination, int *port, int *err)
{
	struct sockaddr_in dest_addr;
	socklen_t dest_addr_len = sizeof(dest_addr);

	printf("fd: %d\n", fd);

	char *dst_str = (char *) malloc(INET_ADDRSTRLEN);
	bzero(dst_str, INET_ADDRSTRLEN);

	if (getsockopt(fd, SOL_IP, SO_ORIGINAL_DST, (struct sockaddr *) &dest_addr, &dest_addr_len)) 
	{
		free(dst_str);
		fprintf(stderr, "getsockopt() failed.");
		*err = errno;
		return -1;
	}

	if(inet_ntop(AF_INET, (void *) &dest_addr.sin_addr, dst_str, INET_ADDRSTRLEN) == NULL)
	{
		free(dst_str);
		fprintf(stderr, "inet_ntop() failed.");
		*err = errno;
		return -1;
	}

	*orig_destination = dst_str;
	*port = (int) ntohs(dest_addr.sin_port);
	return 0;
}

int get_original_destination_6(int fd, char **orig_destination, int *port, int *err) 
{
	struct sockaddr_in6 dest_addr;
	socklen_t dest_addr_len = sizeof(dest_addr);

	printf("fd: %d\n", fd);

	char* dst_str = (char*) malloc(INET6_ADDRSTRLEN);
	bzero(dst_str, INET6_ADDRSTRLEN);

	if (getsockopt(fd, SOL_IPV6, SO_ORIGINAL_DST, (struct sockaddr *) &dest_addr, &dest_addr_len)) {
		free(dst_str);
		fprintf(stderr, "getsockopt() failed.");
		*err = errno;
		return -1;
	}

	if (inet_ntop(AF_INET6, (void *) &dest_addr.sin6_addr, dst_str, INET6_ADDRSTRLEN) == NULL) {
		free(dst_str);
		fprintf(stderr, "inet_ntop() failed.");
		*err = errno;
		return -1;
	}

	*orig_destination = dst_str;
	*port = (int) ntohs(dest_addr.sin6_port);
	return 0;
}