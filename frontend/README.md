# User roles:
1. sekrit-admin
2. client-org-admin
3. client-org-user

# Screens
1. /login
2. /org/admin
3. /org/admin/blueprint
4. /org/admin/users
5. /org/dashboard
6. /org/repos
7. /org/sekrits

# screen in

<!-- ## /sekrit/admin

1. whitelist client-org-admin
 -->

## /Login : client-org-admin

## org/admin
0. 3d party github app/gitlab app integration
	1. add enterprise
		* add multiple org 
	2. add org
2. List of orgs/enterprise
<!-- 3. CRUD user (goto /org/admin/users) -->
3. Invite other users
4. 

<!-- 
## /org/admin/blueprint

1. organize the repos according to line of business
 -->

<!-- ## org/admin/users

1. list of users with their roles and active access token
2. create/update/delete a user
3. user Entity have:
	- role
	- username
	- org
	- LOB : user can belong to multiple LOB (LOB id is coming from org/blueprint)
	- role : admin/SPOC/dev
	-  	 -->

## /org/dashboard

1. Get coverage of onboarded users (# of github commits via sekrit / # of github commit) (90/100) = 90% good
2. show total secrets OPEN (high / med/ low) (background run)
3. show total secrets removed (even by sekrit-client + background run) 
4. show total repos onboarded 
5. overall org health 
6. num. of repos being scanned 
7.  

## /org/repos

1. show list of repos onboarded 
2. show custom regex for the repo
	* add / delete / update regex
	* test regex to check the regex working before adding to repo 	
<!-- 3. add repo (?) -->
4. repo health 

## /org/sekrits

1. list of open secrets in org , 
	* there location in git 
	* severity 
	* verified
	* source (found on git)
	* repo url pointing to exact location of secret
	* commiter email
	* secret source (is this aws secret or jenkins secret or custom or whats the source)
	<!-- * ?line of business (if we have) -->

2. list of closed secrets (secrets which are closed) ()
3. sort by option :	
	- committer
	- severity
	- source
	- repo
	- secret source