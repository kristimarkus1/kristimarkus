export INTERVIEW_NAME="find ./ -type f -name 'Buckingham_Place' -exec cat {} \; | grep 'INTERVIEW' | sed 's/.*#//' "
echo $INTERVIEW_NAME
find ./ -type f -name "interview-$INTERVIEW_NAME" -exec cat {} ;
echo $MAIN_SUSPECT