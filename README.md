# E-Commerce Platform API

### Description
This API supports an e-commerce system for buying products, managing orders, and handling transactions between users with roles such as **Admin**, **Seller**, and **Client**. The business flow includes order statuses such as **pending**, **paid**, **shipping**, **completed**, and **cancel**.

### User Roles
- **Admin**: Has full control over the system, including monitoring orders, changing the order status, and managing users and products.
- **Seller**: Manages the products they sell, processes orders, and updates the order status to **shipping** once payment is confirmed.
- **Client**: The buyer who places orders, makes payments

### Business Flow

1. **Pending**: Orders that have been created but have not been paid yet.
2. **Paid**: Orders that have been successfully paid.
3. **Shipping**: Orders that are being processed and shipped by the seller.
4. **Completed**: Orders that have been successfully delivered and completed.
5. **Cancel**: Orders that have been canceled either by the client or admin.

---

## API Routes

### User Routes

- `POST /user/login`  
  **Description**: User login.  
  **Required Role**: All roles (Client, Seller, Admin)

- `POST /user/register`  
  **Description**: Register a new user.  
  **Required Role**: All roles (Client, Seller, Admin)

- `POST /user/find-by-email`  
  **Description**: Find a user by email.  
  **Required Role**: All roles (Client, Seller, Admin)

- `GET /user/:id`  
  **Description**: Get user data by ID.  
  **Required Role**: All roles (Client, Seller, Admin)  
  **Middleware**: Authentication required.

- `POST /user/find-by-join-date`  
  **Description**: Find users who joined after a certain date.  
  **Required Role**: **Admin**  
  **Middleware**: Authentication required, Admin role required.

### Order Routes

- `POST /order/create`  
  **Description**: Create a new order.  
  **Required Role**: **Client**  
  **Middleware**: Authentication required.

- `POST /order/create-cart-item`  
  **Description**: Add item to the cart.  
  **Required Role**: **Client**  
  **Middleware**: Authentication required.

- `GET /order/cart/:id`  
  **Description**: Get a cart by user ID.  
  **Required Role**: **Client**  
  **Middleware**: Authentication required.

- `POST /order/order`  
  **Description**: Place an order from the cart.  
  **Required Role**: **Client**  
  **Middleware**: Authentication required.

- `POST /order/checkout`  
  **Description**: Complete the checkout and payment process.  
  **Required Role**: **Admin**  
  **Middleware**: Authentication required, Admin role required.

- `POST /order/shipping/:id`  
  **Description**: Mark the order as **shipping**.  
  **Required Role**: **Seller**  
  **Middleware**: Authentication required, Seller role required.

- `POST /order/cancelled/:id`  
  **Description**: Cancel an order.  
  **Required Role**: **Client**  
  **Middleware**: Authentication required.

- `POST /order/completed/:id`  
  **Description**: Mark the order as **completed**.  
  **Required Role**: **Admin**  
  **Middleware**: Authentication required, Admin role required.

- `POST /order/get-five`  
  **Description**: Get the top 5 clients with the highest spending in the last month.  
  **Required Role**: **Admin**  
  **Middleware**: Authentication required, Admin role required.

### Product Routes

- `POST /product/create`  
  **Description**: Add a new product.  
  **Required Role**: **Seller**  
  **Middleware**: Authentication required, Seller role required.

- `GET /product/:id`  
  **Description**: Get a product by ID.  
  **Required Role**: **Client**, **Seller**  
  **Middleware**: Authentication required.

- `GET /product/`  
  **Description**: Get the list of products.  
  **Required Role**: **Client**, **Seller**  
  **Middleware**: Authentication required.

---

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Web Framework**: Gin Gonic
- **ORM**: GORM
- **Middleware**: Custom authentication and authorization
- **Deployment**: Docker (Optional)

## ERD Diagram

![Screenshot from 2025-04-27 12-37-46](https://github.com/user-attachments/assets/40707bcb-088c-477a-903c-1e90caaa4717)


---

## Screenshoot Swagger Open API

![Screenshot from 2025-04-27 12-38-24](https://github.com/user-attachments/assets/e8429249-a76d-491c-b9d9-f6dd76cb7f16)
