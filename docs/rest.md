
## rest api


+ auth
    - POST /auth/signin
    - POST /auth/signout

+ admin
    + account
        - POST /admin/auth/username/registration/response?accept=true
        - POST /admin/auth/username/registration/response?accept=false
    + mileage
        - POST /admin/mileage/username/commission_rate
        - POST /admin/mileage/all/variance/byfile
        - POST /admin/mileage/username/variance/bypost
+ user
    + account
        - POST /user/account/username/registration
        - DELETE /user/account/username/resignation
    + mileage
        - GET /user/mileage/username/current


 




