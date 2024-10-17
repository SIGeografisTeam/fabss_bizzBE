package routes

import (
	"plastiqu_co/controller"
	"plastiqu_co/controller/auth"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the router
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/regis", auth.RegisterUsers).Methods("POST")
	router.HandleFunc("/login", auth.LoginUsers).Methods("POST")

	// User routes (for authenticated users)
	router.HandleFunc("/user/profile", controller.UpdateUserProfile).Methods("PUT")
	router.HandleFunc("/user/password", controller.ChangeUserPassword).Methods("POST")

	// Admin routes (only accessible to admin users)
	// Route untuk admin memperbarui profil pengguna
	router.HandleFunc("/admin/update-user-profile", controller.AdminUpdateUserProfile).Methods("PUT")
	// Route untuk admin memperbarui peran pengguna
	router.HandleFunc("/admin/update-user-role", controller.AdminUpdateUserRole).Methods("PUT")

	// Endpoint untuk kategori
	router.HandleFunc("/categories", controller.CreateCategory).Methods("POST")        // Membuat kategori baru
	router.HandleFunc("/categories", controller.GetCategories).Methods("GET")          // Mengambil semua kategori
	router.HandleFunc("/categories/{id}", controller.GetCategoryByID).Methods("GET")   // Mengambil kategori berdasarkan ID
	router.HandleFunc("/categories/{id}", controller.UpdateCategory).Methods("PUT")    // Mengupdate kategori berdasarkan ID
	router.HandleFunc("/categories/{id}", controller.DeleteCategory).Methods("DELETE") // Menghapus kategori berdasarkan ID

	// produk
	router.HandleFunc("/products", controller.AddProduct).Methods("POST")           // Tambah produk
	router.HandleFunc("/products", controller.GetAllProducts).Methods("GET")        // Dapatkan semua produk
	router.HandleFunc("/products/{id}", controller.GetProductByID).Methods("GET")   // Dapatkan produk berdasarkan ID
	router.HandleFunc("/products/{id}", controller.UpdateProduct).Methods("PUT")    // Update produk berdasarkan ID
	router.HandleFunc("/products/{id}", controller.DeleteProduct).Methods("DELETE") // Hapus produk berdasarkan ID

	// PaymentDetails routes
	router.HandleFunc("/payment_details", controller.CreatePaymentDetails).Methods("POST")
	router.HandleFunc("/payment_details", controller.GetPaymentDetails).Methods("GET")
	router.HandleFunc("/payment_details/{id}", controller.GetPaymentDetailByID).Methods("GET")
	router.HandleFunc("/payment_details/{id}", controller.UpdatePaymentDetails).Methods("PUT")
	router.HandleFunc("/payment_details/{id}", controller.DeletePaymentDetails).Methods("DELETE")

	// Order routes
	router.HandleFunc("/orders", controller.AddOrder).Methods("POST")                 // Menambahkan pesanan baru
	router.HandleFunc("/orders", controller.GetAllOrders).Methods("GET")               // Mendapatkan semua pesanan
	router.HandleFunc("/orders/{id}", controller.GetOrderByID).Methods("GET")          // Mendapatkan pesanan berdasarkan ID
	router.HandleFunc("/orders/{id}", controller.UpdateOrder).Methods("PUT")           // Memperbarui pesanan berdasarkan ID
	router.HandleFunc("/orders/{id}", controller.DeleteOrder).Methods("DELETE")        // Menghapus pesanan berdasarkan ID
	router.HandleFunc("/orders/advance/{id}", controller.AdvanceOrderStatus).Methods("PATCH") // Mengubah status pesanan ke status berikutnya

	// Cart routes
	router.HandleFunc("/carts", controller.AddToCart).Methods("POST")                   // Menambahkan produk ke keranjang
	router.HandleFunc("/carts/{user_id}", controller.GetCartItems).Methods("GET")       // Mendapatkan semua item di keranjang untuk pengguna tertentu
	router.HandleFunc("/carts/{id}", controller.UpdateCartItem).Methods("PUT")          // Memperbarui item di keranjang berdasarkan ID
	router.HandleFunc("/carts/{id}", controller.RemoveFromCart).Methods("DELETE")       // Menghapus item dari keranjang berdasarkan ID
	return router
}
