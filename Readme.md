# Bookings Web Application


- **Built in Go 1.17**
- **PostgresSQL as Database**
- **Boostrap**
## Package
- [chi router](https://github.com/go-chi/chi)
- [alex edwards SCS](https://github.com/alexedwards/scs/v2) as session management
- [nosurf](https://github.com/justinas/nosurf) as middleware, and protect CSRF attack
- [pgx](https://github.com/jackc/pgx/v4) PostgreSQL driver and toolkit for Go
- [simple mail](https://github.com/xhit/go-simple-mail/v2) as Sending Email pkg 
- [Go validator](https://github.com/asaskevich/govalidator)

## Database Diagram
![](https://i.imgur.com/NEZyBX6.png)

## Feature
1. Booking a Room by selecting available day

![](https://i.imgur.com/rzsTs1n.png)


2. Auto Sending confirmation email using goroutine and channel

![](https://i.imgur.com/ksB9O46.png)

3. A Dashboard can Check All Reservation

![](https://i.imgur.com/UsVndkk.png)

4. Edit delete profile on Dashboard

![](https://i.imgur.com/nJGJyc7.png)

5. A Dashboard calendar is straightforward to watch room status (Red one stands for room restriction)

![](https://i.imgur.com/6zgGQLr.png)






