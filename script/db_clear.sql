

/**
mall_message
*/

delete from mall_message.feedback_message;
delete from mall_message.feedback_topic;
delete from mall_message.upload_file;

/**
mall_order
*/
delete from mall_order.cart;
delete from mall_order.order;
delete from mall_order.order_address;
delete from mall_order.order_express;
delete from mall_order.order_pay;
delete from mall_order.order_product;
delete from mall_order.order_refund;
delete from mall_order.order_refund_address;
delete from mall_order.order_refund_image;
delete from mall_order.user_address;

/**
mall_product
*/
delete from mall_product.banner;
delete from mall_product.category;
delete from mall_product.home;
delete from mall_product.home_tags;
delete from mall_product.product;
delete from mall_product.product_html_content;
delete from mall_product.product_image;
delete from mall_product.product_param_basic;
delete from mall_product.product_param_grade;
delete from mall_product.product_sku;
delete from mall_product.product_spec;
delete from mall_product.product_visible;
delete from mall_product.spec;
delete from mall_product.spec_value;
delete from mall_product.upload_file;
delete from mall_product.upload_group;
delete from mall_product.user_favorites;
delete from mall_product.user_history;

/**
mall_user
*/
delete from mall_user.user;
delete from mall_user.merchant;
delete from mall_user.user_grade_log;
delete from mall_user.user_invite_profit_month;
delete from mall_user.user_invite_profit_record;
delete from mall_user.user_register_day;